import type { NextApiRequest, NextApiResponse } from "next";
import * as accountUseCase from "../../../src/usecases";
import * as entities from "../../../src/entities/Account";

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<
    entities.GetAccountControllerOutput | entities.DataError
  >
) {
  switch (req.method) {
    case "GET":
      return getAccountController(req, res);
    default:
      return res
        .status(400)
        .send({ statusCode: 400, message: "Method not allowed" });
  }
}

const getAccountController = async (
  req: NextApiRequest,
  res: NextApiResponse<
    entities.GetAccountControllerOutput | entities.DataError
  >
) => {
  try {
    const { id } = req.query;

    if (id !== undefined) {
      const getAccountControllerInput: entities.GetAccountControllerInput = {
        id: id[0],
      };
    }

    const getAccountUseCaseOutput = await accountUseCase.getAccount(getAccountControllerInput);

    const getAccountControllerOutput: entities.GetAccountControllerOutput =
      {
        items: getAccountUseCaseOutput.map((item) => ({
            id: item.id,
            name: item.name,
            type: item.type,
            kind: item.kind,
        })),
      };

    return res.status(200).send(getAccountControllerOutput);
  } catch (err: any) {
    return res.status(500).json({ statusCode: 500, message: err.message });
  }
};
