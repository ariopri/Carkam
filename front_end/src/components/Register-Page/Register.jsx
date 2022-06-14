import React from 'react';
import './Register.scss';
import { Link } from 'react-router-dom';
import logo from './image/logo-2.png';
import registerImg from './image/register.png';

function Register() {
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
                <h5>
                  Already have an Account?
                  <Link id="Links-signin" to="/Login">
                    Sign In
                  </Link>
                </h5>
              </div>
            </div>
          </div>
        </header>
      </div>
    </>
  );
}

export default Register;
