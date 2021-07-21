import { Router } from "express";
import {
  helloWorld
} from "../controllers/index.controller";

const router: Router = Router();

router.route("/").get(helloWorld);


export default router;
