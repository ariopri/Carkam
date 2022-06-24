import Logo from "./image/KM.png";
import Card from "react-bootstrap/Card";
import { useState } from "react";
import axios from "axios";
import { useParams } from "react-router-dom";
import { useEffect } from "react";

const InfoKampus = () => {
  const baseUrl = "https://reqres.in";
  const token = localStorage.getItem("token");
  if (token === null) {
    console.log("token null");
  } else {
    console.log("token not null");
  }

  const { id } = useParams();

  const [data, setdata] = useState([]);
  const [data2, setdata2] = useState([]);

  const getdata = async () => {
    const res = await axios.get(`${baseUrl}/api/users/${id}`, {
      headers: {
        Authorization: token,
      },
    });
    const res2 = await axios.get(`${baseUrl}/api/users`, {
      headers: {
        Authorization: token,
      },
    });

    setdata(res.data.data);
    setdata2(res2.data.data);
    console.log(res.data.data);
    console.log(res2.data.data);
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
                    {token ? (
                      <button className="btn btn-success  w-100  ">
                        <a
                          className="btn text-light"
                          href={`/Kampus/detail/review/${data.id}`}
                        >
                          <b>R e v i e w</b>
                        </a>
                      </button>
                    ) : (
                      <p className="bg-success text-light p-1  w-100">
                        <center>Lihat review dibawah ini</center>
                      </p>
                    )}

                    <div class="container">
                      {data2.map((item) => {
                        return (
                          <div class="row pt-3">
                            <div class="col-md-12">{item.first_name}</div>
                            <footer className="blockquote-footer">
                              {item.email}
                            </footer>
                          </div>
                        );
                      })}
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
