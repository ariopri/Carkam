import React, { useState } from 'react';
import './Register.scss';
import { Link } from 'react-router-dom';
import logo from './image/logo-2.png';
import registerImg from './image/register.png';

function Register() {
  const [email, setemail] = useState('');
  const [Fusername, setFusername] = useState('');
  const [Susername, setSusername] = useState('');
  const [pwd1, setpwd1] = useState('');
  const [pwd2, setpwd2] = useState('');

  const handlesubmit = (e) => {
    e.preventDefault();
  };

  return (
    <>
      <div className="App">
        <header className="App-header">
          <div className="main-Register">
            <div className="left-side">
              <div className="header">
                <img src={logo} id="logo-img" alt="" srcset="" />
              </div>
              <div className="body">
                <img src={registerImg} id="reg-img" alt="" srcset="" />
              </div>
              <p>“You can teach a student a lesson for a day; but if you can teach him to learn by creating curiosity, he will continue the learning process as long as he lives.”</p>
            </div>
            <div className="right-side">
              <div className="top-right">
                <p>
                  Already have an Account?
                  <Link id="Links-signin" to="/Login">
                    Sign In
                  </Link>
                </p>
              </div>
              <div className="body-right">
                <div className="container">
                  <h1>Create Account!</h1>
                  <form onSubmit={handlesubmit}>
                    <div className="input-group">
                      <h5>First Name</h5>
                      <input
                        type="text"
                        name="Fname"
                        value={Fusername}
                        onChange={(e) => {
                          setFusername(e.target.value);
                        }}
                        id="fname"
                      />
                    </div>
                    <div className="input-group">
                      <h5>Last Name</h5>
                      <input
                        type="text"
                        name="lname"
                        value={Susername}
                        onChange={(e) => {
                          setSusername(e.target.value);
                        }}
                        id="lname"
                      />
                    </div>
                    <div className="input-group">
                      <h5>Email</h5>
                      <input
                        type="Email"
                        name="email"
                        value={email}
                        onChange={(e) => {
                          setemail(e.target.value);
                        }}
                        id="email1"
                      />
                    </div>
                    <div className="input-group">
                      <h5>Password</h5>
                      <input
                        type="password"
                        name="pwd"
                        value={pwd1}
                        onChange={(e) => {
                          setpwd1(e.target.value);
                        }}
                        id="pwd1"
                      />
                    </div>
                    <div className="input-group">
                      <h5>Confirm Password</h5>
                      <input
                        type="password"
                        name="pwd"
                        value={pwd2}
                        onChange={(e) => {
                          setpwd2(e.target.value);
                        }}
                        id="pwd2"
                      />
                    </div>
                    <input type="submit" id="sbtn" value="SUBMIT" />
                  </form>
                </div>
              </div>
            </div>
          </div>
        </header>
      </div>
    </>
  );
}

export default Register;
