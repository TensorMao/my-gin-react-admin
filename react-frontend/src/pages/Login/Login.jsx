import React from 'react'
import {Button, Form, Input, message} from "antd";
import './Login.scss'
import Logo from '../../assets/images/logo512.png'
import axios from "axios";
import {useNavigate} from "react-router-dom";
import storageUtils from "../../utils/storageUtils";
import memoryUtils from "../../utils/memoryUtils";
import {reqLogin} from "../../api";



function Login (){
    const navigate=useNavigate()
    const onFinish = async  (values) => {
        try {
            const res=await reqLogin(values.mobile,values.password)
            if(res.status===200){
                const user=res.data
               //保存用户
                storageUtils.saveUser(user)
                memoryUtils.user=user

                navigate('/home')
                message.success("Succeeded to login.")
           }
             }catch (error) {
            message.error('Failed to login. Please check your mobile and password.');
        }
    };
    const onFinishFailed = (errorInfo) => {
        console.log('Failed:', errorInfo);
    };

        return (
            <div className='login'>
               <div className='login-header'>
                   <img className='logo' src={Logo} alt='logo' />
                   <h1 className="login-h1">Background Management System</h1>
               </div>
                <div className='login-content'>
                    <Form  className='login-form'
                    onFinish={onFinish}
                    onFinishFailed={onFinishFailed}>
                        <Form.Item
                        label="mobile"
                        name="mobile"
                        rules={[
                            { required: true,  message: 'Please input your mobile!' }
                        ]}
                        >

                                <Input placeholder="mobile"/>
                        </Form.Item>
                        <Form.Item
                        label='password'
                        name='password'
                        rules={[{required: true,  message: 'Please input your password!'}]}
                        >

                                <Input.Password placeholder="password"/>



                        </Form.Item>
                        <Form.Item>
                            <Button type="primary" htmlType="submit"
                                    className="login-form-button" style={{ width: '100%'
                            }}>
                                Login
                            </Button>
                        </Form.Item>

                    </Form>


                </div>
            </div>
        )

}
export default Login;