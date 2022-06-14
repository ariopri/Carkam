import './Register.scss';
import { Link } from 'react-router-dom';

const Register = () => {
  return (
    <>
      <div className="App">
        <header className="App-header">
          <div>Register Page</div>
          <Link to="/Login">Move to Login Page</Link>
        </header>
      </div>
    </>
  );
};

export default Register;
