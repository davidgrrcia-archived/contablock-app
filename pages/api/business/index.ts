import type { NextApiRequest, NextApiResponse } from "next";
import * as entities from "../../../src/entities/Business";
import BusinessUseCase from "../../../src/usecases/business.usecase";

const businessUseCase = new BusinessUseCase();

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<
    | entities.ListBusinessControllerOutput
    | entities.DataError
  >
) {
  switch (req.method) {
    case "GET":
      return listBusinessController(req, res);
    default:
      return res
        .status(400)
        .send({ statusCode: 400, message: "Method not allowed" });
  }
}

const listBusinessController = async (
  req: NextApiRequest,
  res: NextApiResponse<
    entities.ListBusinessControllerOutput | entities.DataError
  >
) => {
  try {
    const listBusinessUseCaseOutput =
      await businessUseCase.ListBusinessUseCase();

    const listBusinessControllerOutput: entities.ListBusinessControllerOutput =
        {
          items: listBusinessUseCaseOutput.items.map((item) => ({
            id: item.id,
            name: item.name,
            socialDenomination: item.socialDenomination,
            cif: item.cif,
            logo: item.logo,
          })),
        };

    return res.status(200).send(listBusinessControllerOutput);
  } catch (err: any) {
    console.log(err);
    return res.status(500).json({ statusCode: 500, message: err.message });
  }
};
