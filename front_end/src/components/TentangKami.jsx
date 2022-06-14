import { Container } from "react-bootstrap";
import SlideShow from "./TentangKamiItem/SlideShow";
import Deskripsi from "./TentangKamiItem/Deskripsi";
import OurTeam from "./TentangKamiItem/OurTeam";
import Visimisi from "./TentangKamiItem/Visimisi";

const TentangKami = () => {
  return (
    <>
      <Container>
        <SlideShow />
        <Deskripsi />
        <OurTeam />
        <Visimisi />
      </Container>
    </>
  );
};

export default TentangKami;
