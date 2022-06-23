import Card from "react-bootstrap/Card";
import ListGroup from "react-bootstrap/ListGroup";
import KM from "./image/KM.png";
import axios from "axios";
import { useEffect } from "react";
import { useState } from "react";

const ListKampus = (props) => {
  const dapat = props.cari;
  const baseUrl = "https://reqres.in";
  const token = localStorage.getItem("token");

  const [data, setdata] = useState([]);

  const getdata = async () => {
    const res = await axios.get(`${baseUrl}/api/users`, {
      headers: {
        Authorization: token,
      },
    });
    setdata(res.data.data);
  };
  useEffect(() => {
    getdata();
  }, []);

  return (
    <>
      <div>
        <center>
          <h3>Pilih sepuasnya dan bedah hingga dapat</h3>
        </center>
        <div class="container">
          <div class="row pt-3 ">
            {data.map((item) => {
              if (item.first_name.toLowerCase().includes(dapat)) {
                return (
                  <div class="col-md-3 pb-3 ">
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
                        <Card.Link
                          href={"/Kampus/detail/`}" + item.id}
                          className="btn btn-success"
                        >
                          Lihat
                        </Card.Link>
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
