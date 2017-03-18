import React,{Component} from 'react';
import connectToStores from 'alt-utils/lib/connectToStores';
import TestStore from 'stores/TestStore.js';
import Layout from 'components/common/Layout';

class Test extends Component {

	static getStores(){
		return [TestStore]
	}

	static getPropsFromStores(){
		let state = TestStore.getState();
		console.log('Global State :',state)
		return {
			state : state.testState
		}
	}

	render(){
		return(
			<Layout>
				{this.props.state}
			</Layout>
		)
	}
}

export default connectToStores(Test);
