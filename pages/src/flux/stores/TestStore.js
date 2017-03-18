import Alt from 'bases/Alt.js';
import TestAction from 'actions/TestAction.js';

class TestStore{
	constructor(){
		this.bindListeners({
			handleInitTest: TestAction.initTest
		});
		this.state={
			testState:'hello test'
		}
	}

	handleInitTest = (value)=>{
		console.log('test store value:',value)
	}
}

export default Alt.createStore(TestStore,'TestStore');
