import { Request, Response } from "express";
import fs from "fs/promises";
import path from "path";

export function helloWorld(req: Request, res: Response) {
  res.json({ msg: "Hello World!" });
}

export async function getAvatars(req: Request, res: Response) {
  try {
    const data = await fs.readdir(
      path.join(__dirname, "../public/free_avatars")
    );
    let urls = data.map((file) => `/free_avatars/${file}`);
    res.status(200).json(urls);
  } catch (e) {
    console.log(e);
    return res.status(500).json({ status: "error", msg: e });
  }
}
