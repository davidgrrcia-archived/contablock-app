import type { NextApiRequest, NextApiResponse } from "next";
import * as entities from "../../../src/entities/Business";
import BusinessUseCase from "../../../src/usecases/business.usecase";

const businessUseCase = new BusinessUseCase();

export default async function handler(
    req: NextApiRequest,
    res: NextApiResponse<
        | entities.ListBusinessControllerOutput
        | entities.CreateBusinessControllerOutput
        | entities.DataError
        >
) {
    switch (req.method) {
        case "GET":
            return listBusinessController(req, res);
        case "POST":
            return createBusinessController(req, res);
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

        return res.status(200).send(listBusinessControllerOutput);
    } catch (err: any) {
        console.log(err);
        return res.status(500).json({ statusCode: 500, message: err.message });
    }
};

const createBusinessController = async (
    req: NextApiRequest,
    res: NextApiResponse<
        entities.CreateBusinessControllerOutput | entities.DataError
        >
) => {
    try {
        const createBusinessControllerInput: entities.CreateBusinessUseCaseInput = {
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

        const createBusinessUseCaseOutput =
            await businessUseCase.CreateBusinessUseCase(
                createBusinessControllerInput
            );

        const createBusinessControllerOutput: entities.CreateBusinessControllerOutput =
            {
                id: createBusinessUseCaseOutput.id,
                name: createBusinessUseCaseOutput.name,
                shortDescription: createBusinessUseCaseOutput.shortDescription,
                longDescription: createBusinessUseCaseOutput.longDescription,
                logo: createBusinessUseCaseOutput.logo,
                stockPrice: createBusinessUseCaseOutput.stockPrice,
                urlWebsite: createBusinessUseCaseOutput.urlWebsite,
                ceoName: createBusinessUseCaseOutput.ceoName,
                numberOfEmployees: createBusinessUseCaseOutput.numberOfEmployees,
                country: createBusinessUseCaseOutput.country,
                nif: createBusinessUseCaseOutput.nif,
                fiscalName: createBusinessUseCaseOutput.fiscalName,
                address: createBusinessUseCaseOutput.address,
                phone: createBusinessUseCaseOutput.phone,
                email: createBusinessUseCaseOutput.email,
                state: createBusinessUseCaseOutput.state,
                creationDate: createBusinessUseCaseOutput.creationDate,
                lastUpdate: createBusinessUseCaseOutput.lastUpdate,
            };

        return res.status(201).send(createBusinessControllerOutput);
    } catch (err: any) {
        console.log(err);
        return res.status(500).json({ statusCode: 500, message: err.message });
    }
};
