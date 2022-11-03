import type { NextApiRequest, NextApiResponse } from "next";
import * as entities from "../../../src/entities/Entry";
import EntryUseCase from "../../../src/usecases/entry.usecase";

const entryUseCase = new EntryUseCase();

export default async function handler(
    req: NextApiRequest,
    res: NextApiResponse<
        entities.GetEntryControllerOutput | entities.DataError
        >
) {
  switch (req.method) {
    case "GET":
      return getEntryController(req, res);
    case "PUT":
      return updateEntryController(req, res);
    case "DELETE":
      return deleteEntryController(req, res);
    default:
      return res
          .status(400)
          .send({ statusCode: 400, message: "Method not allowed" });
  }
}

const getEntryController = async (
    req: NextApiRequest,
    res: NextApiResponse<
        entities.GetEntryControllerOutput | entities.DataError
        >
) => {
  try {
    const entryId: string = req.query.entryId
        ? (req.query.entryId as string)
        : "";

    const getEntryUseCaseInput: entities.GetEntryUseCaseInput = {
      id: entryId,
    };

    const getEntryUseCaseOutput = await entryUseCase.GetEntryUseCase(
        getEntryUseCaseInput
    );

    const getEntryControllerOutput: entities.GetEntryControllerOutput = {
      id: getEntryUseCaseOutput.id,
      name: getEntryUseCaseOutput.name,
      shortDescription: getEntryUseCaseOutput.shortDescription,
      longDescription: getEntryUseCaseOutput.longDescription,
      logo: getEntryUseCaseOutput.logo,
      stockPrice: getEntryUseCaseOutput.stockPrice,
      urlWebsite: getEntryUseCaseOutput.urlWebsite,
      ceoName: getEntryUseCaseOutput.ceoName,
      numberOfEmployees: getEntryUseCaseOutput.numberOfEmployees,
      country: getEntryUseCaseOutput.country,
      nif: getEntryUseCaseOutput.nif,
      fiscalName: getEntryUseCaseOutput.fiscalName,
      address: getEntryUseCaseOutput.address,
      phone: getEntryUseCaseOutput.phone,
      email: getEntryUseCaseOutput.email,
      state: getEntryUseCaseOutput.state,
      creationDate: getEntryUseCaseOutput.creationDate,
      lastUpdate: getEntryUseCaseOutput.lastUpdate,
    };

    return res.status(200).send(getEntryControllerOutput);
  } catch (err: any) {
    console.log(err);
    return res.status(500).json({ statusCode: 500, message: err.message });
  }
};

const updateEntryController = async (
    req: NextApiRequest,
    res: NextApiResponse<
        entities.UpdateEntryControllerOutput | entities.DataError
        >
) => {
  try {
    const id: string = req.query.id ? (req.query.id as string) : "";

    const updateEntryUseCaseInput: entities.UpdateEntryUseCaseInput = {
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

    const updateEntryUseCaseOutput =
        await entryUseCase.UpdateEntryUseCase(updateEntryUseCaseInput);

    const updateEntryControllerOutput: entities.UpdateEntryControllerOutput =
        {
          id: updateEntryUseCaseOutput.id,
          name: updateEntryUseCaseOutput.name,
          shortDescription: updateEntryUseCaseOutput.shortDescription,
          longDescription: updateEntryUseCaseOutput.longDescription,
          logo: updateEntryUseCaseOutput.logo,
          stockPrice: updateEntryUseCaseOutput.stockPrice,
          urlWebsite: updateEntryUseCaseOutput.urlWebsite,
          ceoName: updateEntryUseCaseOutput.ceoName,
          numberOfEmployees: updateEntryUseCaseOutput.numberOfEmployees,
          country: updateEntryUseCaseOutput.country,
          nif: updateEntryUseCaseOutput.nif,
          fiscalName: updateEntryUseCaseOutput.fiscalName,
          address: updateEntryUseCaseOutput.address,
          phone: updateEntryUseCaseOutput.phone,
          email: updateEntryUseCaseOutput.email,
          state: updateEntryUseCaseOutput.state,
          creationDate: updateEntryUseCaseOutput.creationDate,
          lastUpdate: updateEntryUseCaseOutput.lastUpdate,
        };

    return res.status(200).send(updateEntryControllerOutput);
  } catch (err: any) {
    console.log(err);
    return res.status(500).json({ statusCode: 500, message: err.message });
  }
};

const deleteEntryController = async (
    req: NextApiRequest,
    res: NextApiResponse<entities.DeleteEntryControllerOutput | entities.DataError >
) => {
  try {
    const id: string = req.query.id ? (req.query.id as string) : "";

    const deleteEntryUseCaseInput: entities.DeleteEntryUseCaseInput = {
      id: id,
    };

    const deleteEntryUseCaseOutput =
        await entryUseCase.DeleteEntryUseCase(deleteEntryUseCaseInput);

    const deleteEntryControllerOutput: entities.DeleteEntryControllerOutput = {
      id: deleteEntryUseCaseOutput.id,
      name: deleteEntryUseCaseOutput.name,
      shortDescription: deleteEntryUseCaseOutput.shortDescription,
      longDescription: deleteEntryUseCaseOutput.longDescription,
      logo: deleteEntryUseCaseOutput.logo,
      stockPrice: deleteEntryUseCaseOutput.stockPrice,
      urlWebsite: deleteEntryUseCaseOutput.urlWebsite,
      ceoName: deleteEntryUseCaseOutput.ceoName,
      numberOfEmployees: deleteEntryUseCaseOutput.numberOfEmployees,
      country: deleteEntryUseCaseOutput.country,
      nif: deleteEntryUseCaseOutput.nif,
      fiscalName: deleteEntryUseCaseOutput.fiscalName,
      address: deleteEntryUseCaseOutput.address,
      phone: deleteEntryUseCaseOutput.phone,
      email: deleteEntryUseCaseOutput.email,
      state: deleteEntryUseCaseOutput.state,
      creationDate: deleteEntryUseCaseOutput.creationDate,
      lastUpdate: deleteEntryUseCaseOutput.lastUpdate,
    };

    return res.status(200).send(deleteEntryControllerOutput);
  } catch (err: any) {
    console.log(err);
    return res.status(500).json({ statusCode: 500, message: err.message });
  }
}
