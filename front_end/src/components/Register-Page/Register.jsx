import './Register.scss';
import { Link } from 'react-router-dom';

const Register = () => {
  return (
    <>
      <div>Register Page</div>
      <Link to="/Login">Move to Login Page</Link>
    </>
  );
};

export default Register;
