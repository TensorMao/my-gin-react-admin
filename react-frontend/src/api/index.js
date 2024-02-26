// 获取所有⻆⾊列表
import ajax from "./ajax";
const BASE = '';


export const reqLogin=(mobile,password)=>{
    return ajax.post(
        BASE+'/auth/login',
        {
            mobile,
        password
        }

    )
}


export const reqLogout=(token)=>{
    return ajax.post(
        BASE+'/auth/logout',
        {},
        {
            timeout:5000,
           headers:{
            Authorization:`Bearer ${token}`
        }
        }
        )
}

export const reqRoles=()=>{
    

}
// 添加⻆⾊
export const reqAddRoles=()=>{}
// 修改权限
export const reqAddAuth=(role)=>{}

export const reqRemoveUser=(user)=>{}

export const reqUsers=()=>{}

