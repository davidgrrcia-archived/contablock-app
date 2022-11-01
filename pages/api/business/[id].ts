import type { NextApiRequest, NextApiResponse } from "next";
import * as entities from "../../../src/entities/Business";
import BusinessUseCase from "../../../src/usecases/business.usecase";

const businessUseCase = new BusinessUseCase();

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<
    entities.GetBusinessControllerOutput | entities.DataError
  >
) {
  switch (req.method) {
    case "GET":
      return getBusinessController(req, res);
    default:
      return res
        .status(400)
        .send({ statusCode: 400, message: "Method not allowed" });
  }
}

const getBusinessController = async (
  req: NextApiRequest,
  res: NextApiResponse<
    entities.GetBusinessControllerOutput | entities.DataError
  >
) => {
  try {
    const businessId: string = req.query.businessId
      ? (req.query.businessId as string)
      : "";

    const getBusinessUseCaseInput: entities.GetBusinessUseCaseInput = {
      id: businessId,
    };

    const getBusinessUseCaseOutput = await businessUseCase.GetBusinessUseCase(
      getBusinessUseCaseInput
    );

    const getBusinessControllerOutput: entities.GetBusinessControllerOutput = {
      id: getBusinessUseCaseOutput.id,
      name: getBusinessUseCaseOutput.name,
      socialDenomination: getBusinessUseCaseOutput.socialDenomination,
      cif: getBusinessUseCaseOutput.cif,
      logo: getBusinessUseCaseOutput.logo,
    };

    return res.status(200).send(getBusinessControllerOutput);
  } catch (err: any) {
    console.log(err);
    return res.status(500).json({ statusCode: 500, message: err.message });
  }
};
