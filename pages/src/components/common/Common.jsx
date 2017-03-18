import React,{ Component } from 'react'
import styles from 'components/common/Common.less';
import { Row, Col, Button, Modal } from 'antd';
import { fenToYuan } from 'services/functions';

//通用水平分割线
export class HR extends React.Component {
    render(){
        return (
            <Row >
                <Col>
                    <hr className={styles.hr}/>
                </Col>
            </Row>
        )
    }
}

//通用全宽度水平分割线
export class FullWidthHR extends React.Component {
    render(){
        return (
            <Row >
                <Col>
                    <hr className={styles.fullHr}/>
                </Col>
            </Row>
        )
    }
}

