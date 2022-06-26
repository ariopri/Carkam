import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import "./Login.scss";
import { Link } from "react-router-dom";
import logo from "./image/logo-carkam.png";
import welcomeImg from "./image/login.png";
import axios from "axios";

function Login() {
  const navigate = useNavigate();
  const baseUrl = "https://reqres.in";
  const [username, setusername] = useState("");
  const [password, setpassword] = useState("");
  const [isError, setisError] = useState(false);
  const [error, seterror] = useState("");
  const email = username;
  const userlogin = async () => {
    const user = {
      email,
      password,
    };
    try {
      const res = await axios.post(`${baseUrl}/api/login`, user);
      localStorage.setItem("token", res.data.token);
      setusername("");
      setpassword("");
      navigate("/Kampus");
      console.log(res.data);
    } catch (err) {
      seterror(err.response.data.message);
      console.log(error);
      setisError(true);
    }
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
                <form>
                  <label>Username</label>
                  <input
                    placeholder="Enter your username"
                    type="text"
                    onChange={(e) => {
                      setusername(e.target.value);
                    }}
                  />
                  <label>Password</label>
                  <input
                    placeholder="Enter your password"
                    type="password"
                    onChange={(e) => {
                      setpassword(e.target.value);
                    }}
                  />
                  <button
                    type="button"
                    id="sub_butt"
                    onClick={() => {
                      userlogin();
                    }}
                  >
                    Login
                  </button>
                  {isError && (
                    <div className="error">
                      <br />
                      <p>Mohon diperiksa kembali yakkk</p>
                    </div>
                  )}
                </form>
                <div className="footer">
                  <h6>
                    Don't have an Account?{" "}
                    <Link to="/Register">Register Now</Link>
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
