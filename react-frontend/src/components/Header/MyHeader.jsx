import {Button, Layout, message, Modal} from "antd";

import "./MyHeader.scss"
import {useLocation, useNavigate} from "react-router-dom";

import storageUtils from "../../utils/storageUtils";
import memoryUtils from "../../utils/memoryUtils";
import {findLabelByKey, menuList} from "../LeftNav/menConfig";
import {reqLogout} from "../../api";

export default function MyHeader(){
    const {Header}=Layout
    const navigate=useNavigate()
    const location = useLocation();
    const currentPath = location.pathname;
    const { confirm } = Modal;
    const user=memoryUtils.user
   //登出
    const logout=() => {
        confirm({
            title:'Are you sure to logout?',
            content:'',
            onOk:async () => {
                //remove token
                const token = storageUtils.getUserToken();
                const res = await reqLogout(token)
                if (res.status === 200) {

                    storageUtils.removeUser()
                    memoryUtils.user={}
                    //导航
                    navigate("/auth/login");
                    message.success("Succeeded to logout.");
                }
            },
            onCancel:()=>{
                console.log('Cancel');
            }


        }
        )

    }
  //获取标题
    function getLabel(){
     const label=findLabelByKey(menuList,currentPath)
        console.log(currentPath+":"+ label)
      return label
    }


    return(<Header
        style={{background: 'white',padding: 0}}
    >
        <div className="userInfo">
            <h3>{getLabel()}</h3>
           <div className='right'>
               <span className="welcome">welcome, {user.data.username}</span>
               <Button className="button" onClick={logout}>Logout</Button>
           </div>

        </div>
    </Header>)


}


