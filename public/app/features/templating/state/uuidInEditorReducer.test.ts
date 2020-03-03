import { toVariablePayload } from './actions';
import { emptyUuid } from './types';
import { reducerTester } from '../../../../test/core/redux/reducerTester';
import { changeToEditorEditMode, changeToEditorListMode, uuidInEditorReducer } from './uuidInEditorReducer';

describe('uuidInEditorReducer', () => {
  describe('when changeToEditorEditMode is dispatched', () => {
    it('then state should be correct', () => {
      const initialState: string | null = null;
      const payload = toVariablePayload({ name: emptyUuid, type: 'query' });
      reducerTester<string | null>()
        .givenReducer(uuidInEditorReducer, initialState)
        .whenActionIsDispatched(changeToEditorEditMode(payload))
        .thenStateShouldEqual(emptyUuid);
    });
  });

  describe('when changeToEditorListMode is dispatched', () => {
    it('then state should be correct', () => {
      const initialState: string | null = emptyUuid;
      reducerTester<string | null>()
        .givenReducer(uuidInEditorReducer, initialState)
        .whenActionIsDispatched(changeToEditorListMode())
        .thenStateShouldEqual(null);
    });
  });
});
