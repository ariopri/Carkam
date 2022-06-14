import React, { useState } from 'react';
import './Login.scss';
import { Link } from 'react-router-dom';
import logo from './image/logo-carkam.png';
import welcomeImg from './image/login.png';

function Login() {
  const [emailval, setemailval] = useState('');
  const [passval, setpassval] = useState('');
  const handlesubmit = (event) => {
    event.preventDefault();
  };

  return (
    <>
      <div className="App">
        <header className="App-header">
          <div className="main-login">
            <div className="login-contain">
              <div className="left-side">
                <div className="img-class">
                  <img src={logo} id="img-id" alt="" srcSet="" />
                </div>
                <form onSubmit={handlesubmit}>
                  <label for="emil1">Email</label>
                  <input
                    placeholder="Enter your email"
                    type="email"
                    value={emailval}
                    onChange={(e) => {
                      setemailval(e.target.value);
                    }}
                    id="emil1"
                  />
                  <label for="pwd1">Password</label>
                  <input
                    placeholder="Enter password"
                    type="password"
                    value={passval}
                    onChange={(e) => {
                      setpassval(e.target.value);
                    }}
                    id="pwd1"
                  />
                  <button type="submit" id="sub_butt">
                    Login
                  </button>
                </form>
                <div className="footer">
                  <h6>
                    Don't have an Account? <Link to="/Register">Register Now</Link>
                  </h6>
                </div>
              </div>
              <div className="right-side">
                <div className="welcomeNote">
                  <h3>Welcome Back !</h3>
                </div>
                <div className="welcomeImg">
                  <img src={welcomeImg} id="wel-img-id" alt="" srcSet="" />
                </div>
              </div>
            </div>
          </div>
        </header>
      </div>
    </>
  );
}

export default Login;
