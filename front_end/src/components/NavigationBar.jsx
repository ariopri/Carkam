const NavigationBar = () => {
  return (
    <>
      <nav className="navbar navbar-expand-lg navbar-light bg-dark">
        <div className="container">
          <a className="navbar-brand  text-light" href="/">
            <b>C A R K A M . I D</b>
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
                <a className="nav-link btn btn-warning text-dark" href="/Login">
                  <b> Hi, Ahmad</b>{" "}
                </a>
              </li>
            </ul>
          </div>
        </div>
      </nav>
    </>
  );
};

export default NavigationBar;
