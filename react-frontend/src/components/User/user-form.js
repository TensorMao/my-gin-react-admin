import {Form, Select} from "antd";
import {Component} from "react";

const {Option}=Select;

class UserForm extends Component{


    render(){
    //const {getFieldDecorator}= this.props.form;
    //const{roles,user}=this.props;
    /*const  formItemLayout = {
        labelCol: { span: 4 },
        wrapperCol: { span: 15 }
    };*/

    return(
        <Form>
            <Form.Item label='username'>

            </Form.Item>


        </Form>
    )

    }

}

export default UserForm;
