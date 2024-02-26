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
function findLabelByKey(menuList, currentPath) {
    for (let i = 0; i < menuList.length; i++) {
        const menuItem = menuList[i];
        if (menuItem.key === currentPath) {
            return menuItem.label;
        }
        if (menuItem.children) {
            const foundInChildren = findLabelByKey(menuItem.children, currentPath);
            if (foundInChildren) {
                return foundInChildren;
            }
        }
    }
    return null;
}

export  {menuList ,findLabelByKey};