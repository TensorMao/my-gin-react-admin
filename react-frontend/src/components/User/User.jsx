import {Component} from "react";
import {Button, Card, message, Modal, Table} from "antd";
import {reqRemoveUser, reqUsers} from "../../api";
import UserForm from "./user-form";


const userTestList=[ {
    _id: 1,
    username: 'John',
    mobile: '1234567890',
    create_time: '2024-02-26T08:00:00Z',
    role_id: 'role1'
},
    {
        _id: 2,
        username: 'Alice',
        mobile: '9876543210',
        create_time: '2024-02-25T10:30:00Z',
        role_id: 'role2'
    }]

const roleTestList=[
    { _id: 'role1', name: 'Admin' },
    { _id: 'role2', name: 'User' }
]

export default class User extends Component{


    state={
        visible:false,//是否显示确认框
       users:[],
       roles:[]
    }

    //初始化
    initColumns=()=>{
        this.columns=[
            {
                title:'username',
                dataIndex:'username'
            },
            {
                title:'mobile',
                dataIndex: 'mobile'
            },
            {
                title:'create_time',
                dataIndex: 'create_time',
                render:(create_time)=>{
                    if(create_time){
                        return new Date(create_time).toLocaleString()
                    }
                }
            },
            {
                title:'role',
                dataIndex: 'role_id',
                render: role_id=>this.state.roles.find(role=>role._id===role_id).name
            },
            {
                title: 'operation',
                render: user => (
                    <span>
                        <Button onClick={() => this.showUpdate(user)}>Modify</Button>
                        <Button onClick={() => this.deleteUser(user)}>Delete</Button>
                    </span>
                )
            }


        ]
    }

    //显示修改界面
    showUpdate=(user)=>{
        this.user=user;
        this.setState({
            visible:true
        })

    }

    //删除指定用户
    deleteUser =async (user)=>{
        Modal.confirm({
            title:`Are you sure to delete the user named ${user.username}? `,
            onOk: async ()=>{
                const res=await reqRemoveUser(user._id);
                if(res.status===200){
                    message.success(`succeeded to delete  ${user.username}`);
                    await this.getUsers();
                }else{
                    message.error(res.error)
                }
            }
        })

    }

    getUsers=async ()=>{
        const res=await reqUsers();
        if(res.status===200){
            const{ users,roles}=res.data
            this.setState({
                users: users, roles:roles
            })
            message.success(`succeeded to get all users`);

        }else{
            message.error(res.error)
        }
    }

    addOrUpdateUser=()=>{
        this.setState({})

    }

    componentWillMount() {
        this.initColumns();
    }
    componentDidMount() {
       // this.getUsers();
        this.setState({users:userTestList,roles:roleTestList})

    }

    handleCancel=e=>{
        this.user=null;
        this.setState({
            visible:false
        })

    }

    showAdd=()=>{
        this.user=null;
        this.setState({visible:true})
    }

    render(){
       // const createUserBtn=(<Button type='primary' onClick={this.showAdd}>Create User</Button>)
        const { users, roles } = this.state;
        const user = this.user || {}
        return (
            <Card  style={{ width:
                    '100%' }}>
                <Table
                    bordered
                    rowKey='_id'
                    dataSource={users}
                    columns={this.columns}
                    pagination={{pageSize: 10  }}>
                </Table>

                {/*<Modal
                    title={user._id ? 'Modify' : 'Add'}
                    visible={this.state.visible}
                    onOk={this.addOrUpdateUser}
                    onCancel={this.handleCancel}
                >
                </Modal>
                <UserForm
                setForm={form=>this.form=form}
                roles={roles}
                user={user}>
                </UserForm>*/}





            </Card>)
    }


}
