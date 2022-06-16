import Container from "react-bootstrap/esm/Container";
import ListKampus from "./KampusItem/ListKampus";
import Pencarian from "./KampusItem/Pencarian";

const Kampus = () => {
  return (
    <>
      <Container>
        <Pencarian />
        <ListKampus />
      </Container>
    </>
  );
};

export default Kampus;
