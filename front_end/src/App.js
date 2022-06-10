import "./App.css";
import NavigationBar from "./components/NavigationBar";
import "./style/SlideShow.css";
import Beranda from "./components/Beranda";
import Berita from "./components/BerandaItem/Berita";

function App() {
  return (
    <>
      <NavigationBar />
      <Beranda />
      <Berita />
    </>
  );
}

export default App;
