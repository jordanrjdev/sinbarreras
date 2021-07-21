import { Request, Response } from "express";

export function createUser(req: Request, res: Response) {
  res.status(200).json({ msg: "create user" });
}
