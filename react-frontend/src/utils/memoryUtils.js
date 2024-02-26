import storageUtils from "./storageUtils";

const user = storageUtils.getUser()
// eslint-disable-next-line import/no-anonymous-default-export
export default {
    user, // ⽤来存储登陆⽤户的信息, 初始值为local中读取的user
}