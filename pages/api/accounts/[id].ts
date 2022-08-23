import type { NextApiRequest, NextApiResponse } from "next";
import * as entities from "../../../src/entities/Account";
import AccountUseCase from "../../../src/usecases/account.usecase";

const accountUseCase = new AccountUseCase();

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<entities.GetAccountControllerOutput | entities.DataError>
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
  res: NextApiResponse<entities.GetAccountControllerOutput | entities.DataError>
) => {
  try {
    const id: string = req.query.id ? (req.query.id as string) : "";

    const getAccountUseCaseInput: entities.GetAccountUseCaseInput = {
      id: id,
    };

    const getAccountUseCaseOutput = await accountUseCase.GetAccountUseCase(
      getAccountUseCaseInput
    );

    const getAccountControllerOutput: entities.GetAccountControllerOutput = {
      id: getAccountUseCaseOutput.id,
      name: getAccountUseCaseOutput.name,
      type: getAccountUseCaseOutput.type,
      kind: getAccountUseCaseOutput.kind,
    };

    return res.status(200).send(getAccountControllerOutput);
  } catch (err: any) {
    console.log(err);
    return res.status(500).json({ statusCode: 500, message: err.message });
  }
};
