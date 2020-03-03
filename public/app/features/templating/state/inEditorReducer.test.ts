import { toVariablePayload } from './actions';
import { NEW_VARIABLE_NAME } from './types';
import { reducerTester } from '../../../../test/core/redux/reducerTester';
import { changeToEditorEditMode, changeToEditorListMode, inEditorReducer } from './inEditorReducer';

describe('inEditorReducer', () => {
  describe('when changeToEditorEditMode is dispatched', () => {
    it('then state should be correct', () => {
      const initialState: string | null = null;
      const payload = toVariablePayload({ name: NEW_VARIABLE_NAME, type: 'query' });
      reducerTester<string | null>()
        .givenReducer(inEditorReducer, initialState)
        .whenActionIsDispatched(changeToEditorEditMode(payload))
        .thenStateShouldEqual(NEW_VARIABLE_NAME);
    });
  });

  describe('when changeToEditorListMode is dispatched', () => {
    it('then state should be correct', () => {
      const initialState: string | null = NEW_VARIABLE_NAME;
      reducerTester<string | null>()
        .givenReducer(inEditorReducer, initialState)
        .whenActionIsDispatched(changeToEditorListMode())
        .thenStateShouldEqual(null);
    });
  });
});
