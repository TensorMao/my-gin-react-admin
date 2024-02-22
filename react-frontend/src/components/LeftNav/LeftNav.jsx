import React, {useState} from 'react'
import {Layout, Menu} from "antd";

import './LeftNav.css'
import menuList from "./menConfig";
import { useNavigate} from "react-router-dom";

export default function LeftNav (props) {
    const {Sider } = Layout;

    const items = menuList.map(item => {
        const newItem = {
            key: item.key,
            icon: item.icon,
            label: item.label
        };
        // 如果存在 children 属性，则将其添加到新对象中
        if (item.children) {
            newItem.children = item.children;
        }
        return newItem;
    });

    const [collapsed, setCollapsed] = useState(false);
    const [selectedKeys, setSelectedKeys] = useState(['/home']);

    const navigate=useNavigate()
    const onClick=(e)=> {
        setSelectedKeys([e.key])
        navigate(e.key,{replace:true})
    }

    return(

        <Sider
            collapsible
            collapsed={collapsed}
            onCollapse={(value) => setCollapsed(value)}
            width={240}
        >
            <Menu
                theme="dark"
                mode="inline"
                defaultSelectedKeys={['/home']}
                selectedKeys={selectedKeys}
                items={items}
                onClick={onClick}
            />

        </Sider>

)
}

