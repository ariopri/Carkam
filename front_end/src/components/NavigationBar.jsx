import { useEffect, useState } from "react";
import logo from "./Login-Page/image/logo-carkam.png";
const NavigationBar = () => {
  const token = localStorage.getItem("token");
  const [isLogin, setIsLogin] = useState(false);

  const cekkin = async () => {
    if (token) {
      setIsLogin(true);
    }
  };

  useEffect(() => {
    cekkin();
  }, []);

  function logout() {
    localStorage.removeItem("token");
    setIsLogin(false);
  }
  return (
    <>
      <nav className="navbar navbar-expand-lg navbar-light bg-dark">
        <div className="container">
          <a className="navbar-brand  text-light" href="/">
            <img src={logo} alt="" width="150px" />
          </a>
          <button
            className="navbar-toggler"
            type="button"
            data-toggle="collapse"
            data-target="#navbarSupportedContent"
            aria-controls="navbarSupportedContent"
            aria-expanded="false"
            aria-label="Toggle navigation"
          >
            <span className="navbar-toggler-icon"></span>
          </button>
          <div className="collapse navbar-collapse" id="navbarSupportedContent">
            <ul className="navbar-nav mx-auto">
              <li className="nav-item active">
                <a className="nav-link  text-light" href="/">
                  Beranda{" "}
                </a>
              </li>
              <li className="nav-item">
                <a className="nav-link  text-light" href="/Kampus">
                  Kampus
                </a>
              </li>
              <li className="nav-item">
                <a className="nav-link  text-light" href="/Tentangkami">
                  Tentang Kami
                </a>
              </li>
              <li className="nav-item">
                <a className="nav-link text-light" href="/FAQ">
                  FAQ
                </a>
              </li>
            </ul>
            <ul className="navbar-nav ml-auto">
              <li className="nav-item active">
                <button type="button" className="btn btn-dark">
                  <b>
                    {isLogin ? (
                      <a
                        className="nav-link btn btn-warning text-dark"
                        onClick={logout}
                        href="/"
                      >
                        <b>Logout</b>
                      </a>
                    ) : (
                      <a
                        href="/Login"
                        className="nav-link btn btn-warning text-dark"
                      >
                        <b>{isLogin ? "Login" : "Login"}</b>
                      </a>
                    )}
                  </b>
                </button>
              </li>
            </ul>
          </div>
        </div>
      </nav>
    </>
  );
};

export default NavigationBar;
