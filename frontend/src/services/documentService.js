import api from './serviceApi';

const VALIDATE_PATH = '/validator/v1/validate';
const SAVE_PATH = '/validator/v1/save';
const GETALL_PATH = '/validator/v1/getAll';

/*export const countAlerts = async () => {
  const header = { headers: { 'Content-Type': 'application/json' } };
  const result = await api.get(`${ALERT_PATH}/count`, header);
  return result;
};*/

export const validateDocument = async (documentNumber) => {
  const header = { headers: { 'Content-Type': 'application/json' } };
  const result = await api.post(`${VALIDATE_PATH}`, {number: documentNumber}, header);
  return result;
}

export const saveDocument = async (documentNumber, documentType) => {
  const header = { headers: { 'Content-Type': 'application/json' } };
  const request = JSON.stringify({
    number: `${documentNumber}`,
    type: `${documentType}`,
    blacklist: false,
  });
  const result = await api.post(`${SAVE_PATH}`, request, header);
  return result;
}

export const getAllDocuments = async () => {
  const header = { headers: { 'Content-Type': 'application/json' } };
  const result = await api.get(`${GETALL_PATH}`, header);
  return result;
}

export default {
  validateDocument,
  saveDocument,
  getAllDocuments
  };
  