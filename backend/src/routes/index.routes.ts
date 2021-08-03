import { Router } from "express";
import { getAvatars, helloWorld } from "../controllers/index.controller";

const router: Router = Router();

router.route("/").get(helloWorld);

router.route("/avatars").get(getAvatars);

export default router;
