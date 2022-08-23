import type { NextApiRequest, NextApiResponse } from "next";
import * as entities from "../../../../src/entities/Business";
import BusinessUseCase from "../../../../src/usecases/business.usecase";

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
    case "PUT":
      return updateBusinessController(req, res);
    case "DELETE":
      return deleteBusinessController(req, res);
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
      shortDescription: getBusinessUseCaseOutput.shortDescription,
      longDescription: getBusinessUseCaseOutput.longDescription,
      logo: getBusinessUseCaseOutput.logo,
      stockPrice: getBusinessUseCaseOutput.stockPrice,
      urlWebsite: getBusinessUseCaseOutput.urlWebsite,
      ceoName: getBusinessUseCaseOutput.ceoName,
      numberOfEmployees: getBusinessUseCaseOutput.numberOfEmployees,
      country: getBusinessUseCaseOutput.country,
      nif: getBusinessUseCaseOutput.nif,
      fiscalName: getBusinessUseCaseOutput.fiscalName,
      address: getBusinessUseCaseOutput.address,
      phone: getBusinessUseCaseOutput.phone,
      email: getBusinessUseCaseOutput.email,
      state: getBusinessUseCaseOutput.state,
      creationDate: getBusinessUseCaseOutput.creationDate,
      lastUpdate: getBusinessUseCaseOutput.lastUpdate,
    };

    return res.status(200).send(getBusinessControllerOutput);
  } catch (err: any) {
    console.log(err);
    return res.status(500).json({ statusCode: 500, message: err.message });
  }
};

const updateBusinessController = async (
  req: NextApiRequest,
  res: NextApiResponse<
    entities.UpdateBusinessControllerOutput | entities.DataError
  >
) => {
  try {
    const id: string = req.query.id ? (req.query.id as string) : "";

    const updateBusinessUseCaseInput: entities.UpdateBusinessUseCaseInput = {
      id: id,
      name: req.body.name,
      shortDescription: req.body.shortDescription,
      longDescription: req.body.longDescription,
      logo: req.body.logo,
      stockPrice: req.body.stockPrice,
      urlWebsite: req.body.urlWebsite,
      ceoName: req.body.ceoName,
      numberOfEmployees: req.body.numberOfEmployees,
      country: req.body.country,
      nif: req.body.nif,
      fiscalName: req.body.fiscalName,
      address: req.body.address,
      phone: req.body.phone,
      email: req.body.email,
      state: req.body.state,
    };

    const updateBusinessUseCaseOutput =
      await businessUseCase.UpdateBusinessUseCase(updateBusinessUseCaseInput);

    const updateBusinessControllerOutput: entities.UpdateBusinessControllerOutput =
      {
        id: updateBusinessUseCaseOutput.id,
        name: updateBusinessUseCaseOutput.name,
        shortDescription: updateBusinessUseCaseOutput.shortDescription,
        longDescription: updateBusinessUseCaseOutput.longDescription,
        logo: updateBusinessUseCaseOutput.logo,
        stockPrice: updateBusinessUseCaseOutput.stockPrice,
        urlWebsite: updateBusinessUseCaseOutput.urlWebsite,
        ceoName: updateBusinessUseCaseOutput.ceoName,
        numberOfEmployees: updateBusinessUseCaseOutput.numberOfEmployees,
        country: updateBusinessUseCaseOutput.country,
        nif: updateBusinessUseCaseOutput.nif,
        fiscalName: updateBusinessUseCaseOutput.fiscalName,
        address: updateBusinessUseCaseOutput.address,
        phone: updateBusinessUseCaseOutput.phone,
        email: updateBusinessUseCaseOutput.email,
        state: updateBusinessUseCaseOutput.state,
        creationDate: updateBusinessUseCaseOutput.creationDate,
        lastUpdate: updateBusinessUseCaseOutput.lastUpdate,
      };

    return res.status(200).send(updateBusinessControllerOutput);
  } catch (err: any) {
    console.log(err);
    return res.status(500).json({ statusCode: 500, message: err.message });
  }
};

const deleteBusinessController = async (
  req: NextApiRequest,
  res: NextApiResponse<entities.DeleteBusinessControllerOutput | entities.DataError >
) => {
  try {
    const id: string = req.query.id ? (req.query.id as string) : "";

    const deleteBusinessUseCaseInput: entities.DeleteBusinessUseCaseInput = {
      id: id,
    };

    const deleteBusinessUseCaseOutput =
      await businessUseCase.DeleteBusinessUseCase(deleteBusinessUseCaseInput);

    const deleteBusinessControllerOutput: entities.DeleteBusinessControllerOutput = {
      id: deleteBusinessUseCaseOutput.id,
      name: deleteBusinessUseCaseOutput.name,
      shortDescription: deleteBusinessUseCaseOutput.shortDescription,
      longDescription: deleteBusinessUseCaseOutput.longDescription,
      logo: deleteBusinessUseCaseOutput.logo,
      stockPrice: deleteBusinessUseCaseOutput.stockPrice,
      urlWebsite: deleteBusinessUseCaseOutput.urlWebsite,
      ceoName: deleteBusinessUseCaseOutput.ceoName,
      numberOfEmployees: deleteBusinessUseCaseOutput.numberOfEmployees,
      country: deleteBusinessUseCaseOutput.country,
      nif: deleteBusinessUseCaseOutput.nif,
      fiscalName: deleteBusinessUseCaseOutput.fiscalName,
      address: deleteBusinessUseCaseOutput.address,
      phone: deleteBusinessUseCaseOutput.phone,
      email: deleteBusinessUseCaseOutput.email,
      state: deleteBusinessUseCaseOutput.state,
      creationDate: deleteBusinessUseCaseOutput.creationDate,
      lastUpdate: deleteBusinessUseCaseOutput.lastUpdate,
    };

    return res.status(200).send(deleteBusinessControllerOutput);
  } catch (err: any) {
    console.log(err);
    return res.status(500).json({ statusCode: 500, message: err.message });
  }
}
