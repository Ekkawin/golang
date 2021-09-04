import { Table } from "antd";
import React, { useEffect } from "react";
import { Axios } from "../axios/axios";

const dataSource = [
  {
    key: "1",
    name: "Mike",
    age: 32,
    address: "10 Downing Street",
  },
  {
    key: "2",
    name: "John",
    age: 42,
    address: "10 Downing Street",
  },
];

const columns = [
  {
    title: "Name",
    dataIndex: "name",
    key: "name",
  },
  {
    title: "Age",
    dataIndex: "age",
    key: "age",
  },
  {
    title: "Address",
    dataIndex: "address",
    key: "address",
  },
];
export const HomePage = () => {
  useEffect(() => {
    (async () => {
      const movies = await Axios.get("http://localhost:8080/movies");
      console.log(`movie`, movies);
    })();
  }, []);
  return (
    <div className="bg-black text-white h-full p-8">
      <div className="flex justify-center items-center">
        <Table dataSource={dataSource} columns={columns} />
      </div>
    </div>
  );
};
