import type { NextApiRequest, NextApiResponse } from "next";
import * as entities from "../../../../../src/entities/Entry";
import EntryUseCase from "../../../../../src/usecases/entry.usecase";

const entryUseCase = new EntryUseCase();

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<
    | entities.ListEntryControllerOutput
    | entities.CreateEntryControllerOutput
    | entities.DataError
  >
) {
  switch (req.method) {
    case "GET":
      return listEntryController(req, res);
    case "POST":
      return createEntryController(req, res);
    default:
      return res
        .status(400)
        .send({ statusCode: 400, message: "Method not allowed" });
  }
}

const listEntryController = async (
  req: NextApiRequest,
  res: NextApiResponse<
    entities.ListEntryControllerOutput | entities.DataError
  >
) => {
  try {
    const listEntryUseCaseOutput =
      await entryUseCase.ListEntryUseCase();

    const listEntryControllerOutput: entities.ListEntryControllerOutput =
      {
        items: listEntryUseCaseOutput.items.map((item) => ({
          id: item.id,
          name: item.name,
          shortDescription: item.shortDescription,
          longDescription: item.longDescription,
          logo: item.logo,
          stockPrice: item.stockPrice,
          urlWebsite: item.urlWebsite,
          ceoName: item.ceoName,
          numberOfEmployees: item.numberOfEmployees,
          country: item.country,
          nif: item.nif,
          fiscalName: item.fiscalName,
          address: item.address,
          phone: item.phone,
          email: item.email,
          state: item.state,
          creationDate: item.creationDate,
          lastUpdate: item.lastUpdate,
        })),
      };

    return res.status(200).send(listEntryControllerOutput);
  } catch (err: any) {
    console.log(err);
    return res.status(500).json({ statusCode: 500, message: err.message });
  }
};

const createEntryController = async (
  req: NextApiRequest,
  res: NextApiResponse<
    entities.CreateEntryControllerOutput | entities.DataError
  >
) => {
  try {
    const createEntryControllerInput: entities.CreateEntryUseCaseInput = {
      id: req.body.id,
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
    };

    const createEntryUseCaseOutput =
      await entryUseCase.CreateEntryUseCase(
        createEntryControllerInput
      );

    const createEntryControllerOutput: entities.CreateEntryControllerOutput =
      {
        id: createEntryUseCaseOutput.id,
        name: createEntryUseCaseOutput.name,
        shortDescription: createEntryUseCaseOutput.shortDescription,
        longDescription: createEntryUseCaseOutput.longDescription,
        logo: createEntryUseCaseOutput.logo,
        stockPrice: createEntryUseCaseOutput.stockPrice,
        urlWebsite: createEntryUseCaseOutput.urlWebsite,
        ceoName: createEntryUseCaseOutput.ceoName,
        numberOfEmployees: createEntryUseCaseOutput.numberOfEmployees,
        country: createEntryUseCaseOutput.country,
        nif: createEntryUseCaseOutput.nif,
        fiscalName: createEntryUseCaseOutput.fiscalName,
        address: createEntryUseCaseOutput.address,
        phone: createEntryUseCaseOutput.phone,
        email: createEntryUseCaseOutput.email,
        state: createEntryUseCaseOutput.state,
        creationDate: createEntryUseCaseOutput.creationDate,
        lastUpdate: createEntryUseCaseOutput.lastUpdate,
      };

    return res.status(201).send(createEntryControllerOutput);
  } catch (err: any) {
    console.log(err);
    return res.status(500).json({ statusCode: 500, message: err.message });
  }
};
