import React from 'react'
import {Routes, Route} from "react-router-dom";
import Login from './pages/Login/Login';
import Home from "./pages/Home/Home";
import Role from "./components/Role/Role";
import User from "./components/User/User";

export default function App() {
    return (

          <Routes>
            <Route path='auth/login' element={<Login />} />
              <Route path='/home' element={<Home />} >
                  <Route path='admin/role' element={<Role />} />
                  <Route path='admin/user' element={<User />} />
                  <Route path='*' element={<div>404 not found</div>} />
              </Route>

              <Route path='*' element={<div>404 not found</div>} />
          </Routes>
    )
}


