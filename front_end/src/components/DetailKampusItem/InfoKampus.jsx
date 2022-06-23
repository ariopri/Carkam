import Logo from "./image/KM.png";
import Card from "react-bootstrap/Card";
import { Link } from "react-router-dom";
import { useState } from "react";
import axios from "axios";
import { useParams } from "react-router-dom";
import { useEffect } from "react";

const InfoKampus = () => {
  const baseUrl = "https://reqres.in";
  const token = localStorage.getItem("token");

  const { id } = useParams();

  const [data, setdata] = useState([]);

  const getdata = async () => {
    const res = await axios.get(`${baseUrl}/api/users/${id}`, {
      headers: {
        Authorization: token,
      },
    });
    setdata(res.data.data);
    console.log(res.data.data);
  };
  useEffect(() => {
    getdata();
  }, []);
  return (
    <>
      <div className="pt-3">
        <div class="container">
          <div class="row">
            <div class="col-md-6">
              <img src={Logo} alt="" width="400px" />
            </div>
            <div class="col-md-6 pt-5">
              <Card>
                <Card.Header className="bg-success text-light">
                  <b>{data.first_name}</b>
                </Card.Header>
                <Card.Body>
                  <blockquote className="blockquote mb-0">
                    <p>
                      <b>Info singkat kampus</b>
                    </p>
                    <p>
                      {data.first_name} {data.last_name}
                    </p>
                    <br />
                    <p>
                      <b>Jurusan : </b> {data.email}
                    </p>
                    <button className="btn btn-success  w-100  ">
                      <a
                        href="/Kampus/detail/review"
                        className="btn text-light"
                      >
                        <b>R e v i e w</b>
                      </a>
                    </button>
                    <div class="container">
                      <div class="row pt-3">
                        <div class="col-md-12">{data.first_name}</div>
                        <footer className="blockquote-footer">
                          {data.email}
                        </footer>
                      </div>
                    </div>
                  </blockquote>
                </Card.Body>
              </Card>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default InfoKampus;
