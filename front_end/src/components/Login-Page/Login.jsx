import React, { useState } from "react";
import { unstable_HistoryRouter } from "react-router-dom";
import "./Login.scss";
import { Link } from "react-router-dom";
import logo from "./image/logo-carkam.png";
import welcomeImg from "./image/login.png";
import axios from "axios";

function Login() {
  const history = unstable_HistoryRouter;
  const baseUrl = "https://my-udemy-api.herokuapp.com/api/v1";
  const [emailval, setemailval] = useState("");
  const [passval, setpassval] = useState("");
  const email = emailval;
  const password = passval;
  const userlogin = async () => {
    const user = {
      email,
      password,
    };
    try {
      const res = await axios.post(`${baseUrl}/login/signin`, user);
      console.log("suksek");
      localStorage.setItem("token", res.data.token);
      setemailval("");
      setpassval("");
      history.push("/Kampus");
    } catch (err) {
      alert("Ups salah");
      console.log("salah");
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
                  <label>Email</label>
                  <input
                    placeholder="Enter your email"
                    type="email"
                    name={emailval}
                    onChange={(e) => {
                      setemailval(e.target.value);
                    }}
                    id="emil1"
                  />
                  <label>Password</label>
                  <input
                    placeholder="Enter your password"
                    type="password"
                    name={passval}
                    onChange={(e) => {
                      setpassval(e.target.value);
                    }}
                    id="pwd1"
                  />
                  <button
                    type="submit"
                    id="sub_butt"
                    onClick={() => {
                      userlogin();
                    }}
                  >
                    Login
                  </button>
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
