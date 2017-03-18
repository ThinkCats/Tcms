import Alt from 'bases/Alt.js';

class TestAction{
    initTest(){
        console.log("init my test");
        return 'hello';
    }
}

export default Alt.createActions(TestAction);
