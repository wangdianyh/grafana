import React, { Component } from 'react';
import { debounce } from 'lodash';
import { Input } from '@grafana/ui';

export interface Props {
  onChange: (alignmentPeriod: any) => void;
  value: string;
}

export interface State {
  value: string;
}

export class LogsQueryField extends Component<Props, State> {
  propagateOnChange: (value: any) => void;

  constructor(props: Props) {
    super(props);
    this.propagateOnChange = debounce(this.props.onChange, 500);
    this.state = { value: '' };
  }

  componentDidMount() {
    this.setState({ value: this.props.value });
  }

  UNSAFE_componentWillReceiveProps(nextProps: Props) {
    if (nextProps.value !== this.props.value) {
      this.setState({ value: nextProps.value });
    }
  }

  onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    this.setState({ value: e.target.value });
    this.propagateOnChange(e.target.value);
  };

  render() {
    return (
      <>
        <Input type="text" className="gf-form-input" value={this.state.value} onChange={this.onChange} />
      </>
    );
  }
}
