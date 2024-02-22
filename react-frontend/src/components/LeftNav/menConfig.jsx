import {HomeOutlined, UserOutlined} from "@ant-design/icons";


const menuList=[
    {
        label: "Home",
        key: '/home',
        icon: <HomeOutlined />,

    },
    {

        label: 'Admin',
        key: '/home/admin',
        icon: <UserOutlined />  ,
        children:[
            {
                label: 'User Management',
                key: '/home/admin/user',
                icon: <UserOutlined /> ,
            },

            {
                label: 'Role Management',
                key: '/home/admin/role',
                icon: <UserOutlined /> ,
            },

        ]

    },

]

export default menuList