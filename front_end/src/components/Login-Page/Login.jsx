import './Login.scss';
import { Link } from 'react-router-dom';

function Login() {
  const handlesubmit = (event) => {
    event.preventDefault();
  };

  return (
    <>
      <div className="main-login">
        Login Page <br />
        <Link to="/Register">Move to Register Page</Link>
        <div className="login-contain">
          <div className="left-side">
            <form onSubmit={handlesubmit}>
              <label for="emil1">Email</label>
              <input placeholder="Enter your email" type="email" id="emil1" />
              <label for="pwd1">Password</label>
              <input placeholder="Enter password" type="password" id="pwd1" />
              <button type="submit" id="sub_butt">
                Submit
              </button>
            </form>
          </div>
          <div className="right-side"></div>
        </div>
      </div>
    </>
  );
}

export default Login;
