import Card from "react-bootstrap/Card";
import ListGroup from "react-bootstrap/ListGroup";
import KM from "./image/KM.png";
import { useEffect } from "react";
import useStore from "./Store";

const ListKampus = (props) => {
  const dapat = props.cari;

  const data = useStore((state) => state.data);
  const getdata = useStore((state) => state.getdata);

  useEffect(() => {
    getdata();
  }, []);

  return (
    <>
      <div>
        <center>
          <h3>Pilih sepuasnya dan bedah hingga dapat</h3>
        </center>
        <div className="container">
          <div className="row pt-3 ">
            {data.map((item) => {
              if (item.first_name.toLowerCase().includes(dapat)) {
                return (
                  <div className="col-md-3 pb-3 ">
                    <Card style={{ width: "18rem" }}>
                      <Card.Img variant="top" src={KM} />
                      <Card.Body>
                        <Card.Title>{item.first_name}</Card.Title>
                        <Card.Text>
                          <p>
                            {item.first_name} {item.last_name}
                          </p>
                        </Card.Text>
                      </Card.Body>
                      <ListGroup className="list-group-flush">
                        <ListGroup.Item>{item.email}</ListGroup.Item>
                      </ListGroup>
                      <Card.Body>
                        <a
                          className="btn btn-success w-100"
                          href={`/Kampus/detail/${item.id}`}
                        >
                          Lihat
                        </a>
                      </Card.Body>
                    </Card>
                  </div>
                );
              }
            })}
          </div>
        </div>
      </div>
    </>
  );
};
export default ListKampus;
