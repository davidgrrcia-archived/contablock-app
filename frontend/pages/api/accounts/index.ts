import type { NextApiRequest, NextApiResponse } from "next";
import * as entities from "../../../src/entities/Account";
import AccountUseCase from "../../../src/usecases/account.usecase";

const accountUseCase = new AccountUseCase();

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
    const listAccountsUseCaseOutput =
      await accountUseCase.ListAccountsUseCase();

    const listAccountsControllerOutput: entities.ListAccountsControllerOutput =
      {
        items: listAccountsUseCaseOutput.items.map((item) => ({
          id: item.id,
          name: item.name,
          group: item.group,
          category: item.category,
        })),
      };

    return res.status(200).send(listAccountsControllerOutput);
  } catch (err: any) {
    console.log(err);
    return res.status(500).json({ statusCode: 500, message: err.message });
  }
};
