import SlideShow from "./BerandaItem/SlideShow";
import Testimoni from "./BerandaItem/Testimoni";
import Berita from "./BerandaItem/Berita";
// import { useEffect, useState } from "react";
// import Container from "react-bootstrap/Container";
const Beranda = () => {
  // const [data, setData] = useState([]);

  // useEffect(() => {
  //   fetch("https://jsonplaceholder.typicode.com/posts").then((res) =>
  //     res.json().then((data) => setData(data))
  //   );
  // }, []);

  return (
    <>
      <SlideShow />
      <Testimoni />
      <Berita />
      {/* <Container>{JSON.stringify(data)}</Container> */}
    </>
  );
};
export default Beranda;
