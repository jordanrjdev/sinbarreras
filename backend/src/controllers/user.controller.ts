import { query, Request, Response } from "express";
import jwt from "jsonwebtoken";
import { OkPacket, RowDataPacket } from "mysql2";
import { connect } from "../config/database";
import { User } from "../interfaces/types";

export async function createUser(req: Request, res: Response) {
  const user: User = req.body;
  try {
    const conn = await connect();
    await conn.query("set @error = ''");
    await conn.query<RowDataPacket[]>(
      `CALL funsiba.sp_create_user('${user.username}', ${user.avatar_id}, @error);`
    );
    const result = await conn.query<RowDataPacket[]>(`SELECT @error as error;`);
    res.status(200).json(result[0][0]);
  } catch (e) {
    console.log(e);
    return res.status(500).json({ status: "error", msg: e });
  }
}

export async function getAllUsers(req: Request, res: Response) {
  try {
    const conn = await connect();
    const result = await conn.query<RowDataPacket[]>(
      `SELECT * FROM funsiba.user;`
    );
    res.status(200).json(result[0]);
  } catch (e) {
    console.log(e);
    return res.status(500).json({ status: "error", msg: e });
  }
}

export async function getUser(req: Request, res: Response) {
  const user_id = req.params.id;
  try {
    const conn = await connect();
    const result = await conn.query<RowDataPacket[]>(
      "SELECT * FROM funsiba.user WHERE user_id = ?",
      [user_id]
    );
    res.status(200).json(result[0][0]);
  } catch (e) {
    console.log(e);
    return res.status(500).json({ status: "error", msg: e });
  }
}

export async function deleteUser(req: Request, res: Response) {
  const { id } = req.body;
  try {
    const conn = await connect();
    await conn.query("delete from user where user_id = ?", [id]);
    res.status(200).json({ status: "success", msg: "User deleted" });
  } catch (e) {
    console.log(e);
    return res.status(500).json({ status: "error", msg: e });
  }
}

export function login(req: Request, res: Response) {
  const { username } = req.body;

  const token = jwt.sign({ username }, "secret", { expiresIn: "5h" });

  res.status(200).json({ token });
}
