// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from "next";
import * as accountUseCase from "../../../src/usecases";
import * as entities from "../../../src/entities/Account";

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<
    | entities.ListAccountsControllerOutput
    | entities.DataError
  >
) {
  switch (req.method) {
    case "GET":
      return listAccounts(req, res);
    default:
      return res
        .status(400)
        .send({ statusCode: 400, message: "Method not allowed" });
  }
}

const listAccounts = async (
  req: NextApiRequest,
  res: NextApiResponse<
    entities.ListAccountsControllerOutput | entities.DataError
  >
) => {
  try {
    const accounts = await accountUseCase.listAccounts();
    const items: entities.GetAccountControllerOutput[] = accounts.map(
      (account) => ({
        id: account.id,
        name: account.name,
        type: account.type,
        kind: account.kind,
      })
    ); //todo: think if do this in the controller or the usecase

    return res.status(200).send({ items: items });
  } catch (err: any) {
    return res.status(500).json({ statusCode: 500, message: err.message });
  }
};