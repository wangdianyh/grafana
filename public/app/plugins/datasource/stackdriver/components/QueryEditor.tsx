import React, { PureComponent } from 'react';
import _ from 'lodash';
import events from 'app/core/app_events';
import { TimeSeries, SelectableValue, ExploreQueryFieldProps, PanelEvents } from '@grafana/data';
import { Aggregations, Metrics, Filters, GroupBys, Alignments, AlignmentPeriods, AliasBy, Help } from './';
import { StackdriverQuery, MetricDescriptor } from '../types';
import { getAlignmentPickerData, toOption } from '../functions';
import StackdriverDatasource from '../datasource';

interface State extends StackdriverQuery {
  variableOptions: Array<SelectableValue<string>>;
  variableOptionGroup: SelectableValue<string>;
  alignOptions: Array<SelectableValue<string>>;
  lastQuery: string;
  lastQueryError: string;
  labels: any;
  [key: string]: any;
}

export type Props = ExploreQueryFieldProps<StackdriverDatasource, StackdriverQuery>;

export const DefaultTarget: State = {
  defaultProject: 'loading project...',
  metricType: '',
  metricKind: '',
  valueType: '',
  refId: '',
  service: '',
  unit: '',
  crossSeriesReducer: 'REDUCE_MEAN',
  alignmentPeriod: 'stackdriver-auto',
  perSeriesAligner: 'ALIGN_MEAN',
  groupBys: [],
  filters: [],
  filter: [],
  aliasBy: '',
  alignOptions: [],
  lastQuery: '',
  lastQueryError: '',
  usedAlignmentPeriod: '',
  labels: {},
  variableOptionGroup: {},
  variableOptions: [],
};

export class QueryEditor extends PureComponent<Props, State> {
  state: State = DefaultTarget;

  async componentDidMount() {
    const { query, datasource } = this.props;
    events.on(PanelEvents.dataReceived, this.onDataReceived.bind(this));
    events.on(PanelEvents.dataError, this.onDataError.bind(this));
    const { perSeriesAligner, alignOptions } = getAlignmentPickerData(query, datasource.templateSrv);
    const variableOptionGroup = {
      label: 'Template Variables',
      expanded: false,
      options: datasource.variables.map(toOption),
    };
    this.setState({
      ...this.props.query,
      alignOptions,
      perSeriesAligner,
      variableOptionGroup,
      variableOptions: variableOptionGroup.options,
    });

    datasource.getLabels(query.metricType, query.refId, query.groupBys).then(labels => this.setState({ labels }));
  }

  componentWillUnmount() {
    events.off(PanelEvents.dataReceived, this.onDataReceived);
    events.off(PanelEvents.dataError, this.onDataError);
  }

  onDataReceived(dataList: TimeSeries[]) {
    const series = dataList.find((item: any) => item.refId === this.props.query.refId);
    if (series) {
      this.setState({
        lastQuery: decodeURIComponent(series.meta.rawQuery),
        lastQueryError: '',
        usedAlignmentPeriod: series.meta.alignmentPeriod,
      });
    }
  }

  onDataError(err: any) {
    let lastQuery;
    let lastQueryError;
    if (err.data && err.data.error) {
      lastQueryError = this.props.datasource.formatStackdriverError(err);
    } else if (err.data && err.data.results) {
      const queryRes = err.data.results[this.props.query.refId];
      lastQuery = decodeURIComponent(queryRes.meta.rawQuery);
      if (queryRes && queryRes.error) {
        try {
          lastQueryError = JSON.parse(queryRes.error).error.message;
        } catch {
          lastQueryError = queryRes.error;
        }
      }
    }
    this.setState({ lastQuery, lastQueryError });
  }

  onMetricTypeChange = async ({ valueType, metricKind, type, unit }: MetricDescriptor) => {
    const { onChange, onRunQuery, query, datasource } = this.props;
    const { perSeriesAligner, alignOptions } = getAlignmentPickerData(
      { valueType, metricKind, perSeriesAligner: this.state.perSeriesAligner },
      datasource.templateSrv
    );
    const labels = await this.props.datasource.getLabels(type, query.refId, query.groupBys);
    this.setState(
      {
        alignOptions,
        perSeriesAligner,
        metricType: type,
        unit,
        valueType,
        metricKind,
        labels,
      },
      () => {
        onChange(this.state);
        onRunQuery();
      }
    );
  };

  onGroupBysChange(value: string[]) {
    const { query, datasource } = this.props;
    this.setState({ groupBys: value }, () => {
      this.props.onChange(this.state);
      this.props.onRunQuery();
    });
    datasource.getLabels(query.metricType, query.refId, value).then(labels => this.setState({ labels }));
  }

  onPropertyChange(prop: string, value: string[]) {
    this.setState({ [prop]: value }, () => {
      this.props.onChange(this.state);
      this.props.onRunQuery();
    });
  }

  render() {
    const {
      usedAlignmentPeriod,
      defaultProject,
      metricType,
      crossSeriesReducer,
      groupBys,
      filters,
      perSeriesAligner,
      alignOptions,
      alignmentPeriod,
      aliasBy,
      lastQuery,
      lastQueryError,
      labels,
      variableOptionGroup,
      variableOptions,
    } = this.state;
    const { datasource } = this.props;

    return (
      <>
        <Metrics
          templateSrv={datasource.templateSrv}
          defaultProject={defaultProject}
          metricType={metricType}
          templateVariableOptions={variableOptions}
          datasource={datasource}
          onChange={this.onMetricTypeChange}
        >
          {metric => (
            <>
              <Filters
                labels={labels}
                filters={filters}
                onChange={value => this.onPropertyChange('filters', value)}
                variableOptionGroup={variableOptionGroup}
              />
              <GroupBys
                groupBys={Object.keys(labels)}
                values={groupBys}
                onChange={this.onGroupBysChange.bind(this)}
                variableOptionGroup={variableOptionGroup}
              />
              <Aggregations
                metricDescriptor={metric}
                templateVariableOptions={variableOptions}
                crossSeriesReducer={crossSeriesReducer}
                groupBys={groupBys}
                onChange={value => this.onPropertyChange('crossSeriesReducer', value)}
              >
                {displayAdvancedOptions =>
                  displayAdvancedOptions && (
                    <Alignments
                      alignOptions={alignOptions}
                      templateVariableOptions={variableOptions}
                      perSeriesAligner={perSeriesAligner}
                      onChange={value => this.onPropertyChange('perSeriesAligner', value)}
                    />
                  )
                }
              </Aggregations>
              <AlignmentPeriods
                templateSrv={datasource.templateSrv}
                templateVariableOptions={variableOptions}
                alignmentPeriod={alignmentPeriod}
                perSeriesAligner={perSeriesAligner}
                usedAlignmentPeriod={usedAlignmentPeriod}
                onChange={value => this.onPropertyChange('alignmentPeriod', value)}
              />
              <AliasBy value={aliasBy} onChange={value => this.onPropertyChange('aliasBy', value)} />
              <Help datasource={datasource} rawQuery={lastQuery} lastQueryError={lastQueryError} />
            </>
          )}
        </Metrics>
      </>
    );
  }
}
