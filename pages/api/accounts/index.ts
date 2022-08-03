// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from "next";
import * as accountUseCase from "../../../src/usecases";

type AccountOutput = {
  name: string;
};

type ListAccountsOutput = {
  accounts: AccountOutput[];
};

type DataError = {
  statusCode: number;
  message: string;
};

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<ListAccountsOutput | AccountOutput | DataError>
) {
  switch (req.method) {
   
    default:
      return res
        .status(400)
        .send({ statusCode: 400, message: "Method not allowed" });
  }
}

const listAccounts = async (
  req: NextApiRequest,
  res: NextApiResponse<ListAccountsOutput | DataError>
) => {
  try {
    const accounts = await accountUseCase.listAccounts();
    const items: AccountOutput[] = accounts.map((account) => ({
      name: account.name,
    }));

    return res.status(200).send({ accounts: items });
  } catch (err: any) {
    return res.status(500).json({ statusCode: 500, message: err.message });
  }
};
