import React, {useEffect, useState} from 'react'
import {Layout, Modal} from "antd";
import LeftNav from "../../components/LeftNav/LeftNav";
import AppContent from "../../components/Content/Content";
import MyHeader from "../../components/Header/MyHeader";
import {Link, Navigate, useNavigate} from "react-router-dom";
import memoryUtils from "../../utils/memoryUtils";


function Home ()  {
    const navigate=useNavigate()
    const user = memoryUtils.user;
    /*useEffect(() => {
        // 在组件加载完成后检查是否存在 access_token
       if (!user.data|| !user.data.access_token) {
         navigate('/auth/login'); // 如果不存在，则跳转到登录页面
      }
    }, [navigate, user.data]
    ); // 空数组作为第二个参数，表示只在组件加载完成后执行一次*/

    const [visible, setVisible] = useState(true);
    const handleOk=()=> {
        setVisible(false);
        navigate('/auth/login');
    }


    return (
        (!user.data|| !user.data.access_token)?(
                   <Modal
                   title="提示"
                   visible={visible}
                   onOk={handleOk}
               >
                   <p>您还未登录，请先登录！</p>
                   {/* 这里可以添加登录按钮 */}
                   </Modal>


        ):(
            <Layout
                    style={{
                        minHeight: '100vh',
                    }}
                >
                    <LeftNav />
                    <Layout>
                        <MyHeader />
                        <AppContent />
                    </Layout>
                </Layout>
            )



        )
}
export default Home;