import React, {useEffect} from 'react'
import {Layout} from "antd";
import LeftNav from "../../components/LeftNav/LeftNav";
import AppContent from "../../components/Content/Content";
import MyHeader from "../../components/Header/MyHeader";
import storageUtils from "../../utils/storageUtils";
import {useNavigate} from "react-router-dom";

function Home ()  {
    const navigate=useNavigate()
    const user = storageUtils.getUser();
    useEffect(() => {

        // 在组件加载完成后检查是否存在 access_token
       if (!user.data|| !user.data.access_token) {
         navigate('/auth/login'); // 如果不存在，则跳转到登录页面
      }
    }, [navigate, user.data]
    ); // 空数组作为第二个参数，表示只在组件加载完成后执行一次


    return (
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
}
export default Home;