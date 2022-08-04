import type { NextApiRequest, NextApiResponse } from "next";
import * as accountUseCase from "../../../src/usecases";
import * as entities from "../../../src/entities/Account";

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<
    entities.ListAccountsControllerOutput | entities.DataError
  >
) {
  switch (req.method) {
    case "GET":
      return listAccountsController(req, res);
    default:
      return res
        .status(400)
        .send({ statusCode: 400, message: "Method not allowed" });
  }
}

const listAccountsController = async (
  req: NextApiRequest,
  res: NextApiResponse<
    entities.ListAccountsControllerOutput | entities.DataError
  >
) => {
  try {
    const listAccountsUseCaseOutput = await accountUseCase.listAccountsUseCase();

    const listAccountsControllerOutput: entities.ListAccountsControllerOutput =
      {
        items: listAccountsUseCaseOutput.map((item) => ({
            id: item.id,
            name: item.name,
            type: item.type,
            kind: item.kind,
        })),
      };

    return res.status(200).send(listAccountsControllerOutput);
  } catch (err: any) {
    return res.status(500).json({ statusCode: 500, message: err.message });
  }
};
