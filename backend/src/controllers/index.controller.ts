import { Request, Response } from "express";

export function helloWorld(req : Request, res : Response) {
    res.json({msg : "Hello World!"});
}