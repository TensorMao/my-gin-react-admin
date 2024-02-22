import {Button, Layout, message} from "antd";

import "./MyHeader.scss"
import {useNavigate} from "react-router-dom";
import axios from "axios";
import storageUtils from "../../utils/storageUtils";

export default function MyHeader(){
    const {Header}=Layout
    const navigate=useNavigate()
    const user=storageUtils.getUser()



    const logout=async () => {
        console.log("user logout")
        const token=storageUtils.getUserToken();
        const instance=axios.create({
            baseURL:'',
            timeout:5000,
            headers:{
                Authorization:`Bearer ${token}`
            }

        })
        const res = await instance.post("/auth/logout")

        if(res.status===200) {
            storageUtils.removeUser()
            navigate("/auth/login")
            message.success("Succeeded to logout.")
        }
    }

    return(<Header
        style={{
            background: 'white',padding: 0
        }}
    >
        <div className="userInfo">
            <span className="welcome">欢迎, admin</span>
            <Button className="button" onClick={logout}>Logout</Button>
        </div>
    </Header>)


}


