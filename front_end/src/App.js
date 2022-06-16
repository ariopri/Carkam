import './App.css';
import NavigationBar from './components/NavigationBar';
import './style/SlideShow.css';
import Beranda from './components/Beranda';
import Footer from './components/Footer';
import TentangKami from './components/TentangKami';
import Login from './components/Login-Page/Login';
import Register from './components/Register-Page/Register';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
function App() {
  return (
    <>
      {/* <div className="App">
        <header className="App-header"> */}
      <NavigationBar />
      <Router>
        <Routes>
          <Route path="/" element={<Beranda />} />
          <Route path="/Tentangkami" element={<TentangKami />} />
          <Route path="/Login" element={<Login />} />
          <Route path="/Register" element={<Register />} />
        </Routes>
      </Router>
      <Footer />
      {/* </header>
      </div> */}
    </>
  );
}

export default App;
