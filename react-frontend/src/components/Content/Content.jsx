import { Layout} from "antd";
import {Outlet} from "react-router-dom";

function AppContent () {
    const { Content } = Layout;
    return (
        <Content
            style={{
                margin: '24px 16px',
                padding: 24,
                background: '#fff',
                minHeight: 280,
            }}>
           <Outlet />

        </Content>
);
}

export default AppContent;